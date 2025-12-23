import io
from typing import List, Dict
import pandas as pd
from fpdf import FPDF

from domain.loan.loan import Loan
from domain.loan.loan_repo import LoanRepositoryAbc
from api.generated.loan_analytic_pb2_grpc import LoanAnalyticServiceServicer
from api.generated.loan_analytic_pb2 import *


def _create_pdf_report(title: str, data: Dict[str, str]) -> bytes:
    """Создает простой PDF отчет."""
    pdf = FPDF()
    pdf.add_page()
    pdf.set_font("Helvetica", "B", 16)
    pdf.cell(0, 10, title, 1, 1, 'C')
    pdf.set_font("Helvetica", "", 12)
    for key, value in data.items():
        pdf.cell(0, 10, f"{key}: {value}", 0, 1)
    return pdf.output(dest='S').encode('latin-1')


def _create_excel_report(data_df: pd.DataFrame) -> bytes:
    """Создает Excel отчет из DataFrame."""
    output = io.BytesIO()
    with pd.ExcelWriter(output, engine='openpyxl') as writer:
        data_df.to_excel(writer, index=False, sheet_name='Report')
    return output.getvalue()


class LoanAnalyticService(LoanAnalyticServiceServicer):
    def __init__(self, loan_repository: LoanRepositoryAbc):
        self.loan_repository = loan_repository

    async def GetTotalActiveLoanPortfolio(self, request: GetTotalActiveLoanPortfolioRequest, context):
        loans: List[Loan] = await self.loan_repository.get_all()
        df = pd.DataFrame(l.__dict__ for l in loans)
        df_active = df[df['status'] == True]
        total_sum = df_active['principal'].sum()

        if request.format == ReportFormat.PDF:
            pdf_content = _create_pdf_report(
                title="Total Active Loan Portfolio",
                data={"Total Portfolio": f"${total_sum:,.2f}"}
            )
            file_response = FileResponse(
                filename="total_portfolio.pdf",
                content=pdf_content,
                mime_type="application/pdf"
            )
            return TotalActiveLoanPortfolioResponse(file_response=file_response)

        if request.format == ReportFormat.EXCEL:
            report_df = pd.DataFrame([{"Total Active Portfolio": total_sum}])
            excel_content = _create_excel_report(report_df)
            file_response = FileResponse(
                filename="total_portfolio.xlsx",
                content=excel_content,
                mime_type="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
            )
            return TotalActiveLoanPortfolioResponse(file_response=file_response)

        # По умолчанию возвращаем структурированный ответ
        structured_response = TotalActiveLoanPortfolioResponse.StructuredResponse(total_portfolio=str(total_sum))
        return TotalActiveLoanPortfolioResponse(structured_response=structured_response)

    async def GetAverageLoanInterestRate(self, request: GetAverageLoanInterestRateRequest, context):
        loans: List[Loan] = await self.loan_repository.get_all()
        df = pd.DataFrame(l.__dict__ for l in loans)
        df_active = df[df['status'] == True]

        if df_active.empty:
            weighted_average_str = "0.00"
        else:
            df_active['weighted_interest'] = df_active['principal'] * df_active['interest_rate']
            weighted_average = df_active['weighted_interest'].sum() / df_active['principal'].sum()
            weighted_average_str = f"{weighted_average:.4f}"

        if request.format == ReportFormat.PDF:
            pdf_content = _create_pdf_report(
                title="Average Loan Interest Rate (Weighted)",
                data={"Average Rate": weighted_average_str}
            )
            file_response = FileResponse(filename="average_rate.pdf", content=pdf_content, mime_type="application/pdf")
            return AverageLoanInterestRateResponse(file_response=file_response)

        if request.format == ReportFormat.EXCEL:
            report_df = pd.DataFrame([{"Average Weighted Interest Rate": weighted_average_str}])
            excel_content = _create_excel_report(report_df)
            file_response = FileResponse(
                filename="average_rate.xlsx",
                content=excel_content,
                mime_type="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
            )
            return AverageLoanInterestRateResponse(file_response=file_response)

        structured_response = AverageLoanInterestRateResponse.StructuredResponse(average_interest_rate=weighted_average_str)
        return AverageLoanInterestRateResponse(structured_response=structured_response)

    async def GetLoanDistributionByStatus(self, request: GetLoanDistributionByStatusRequest, context):
        loans: List[Loan] = await self.loan_repository.get_all()
        df = pd.DataFrame(l.__dict__ for l in loans)
        status_counts = df['status'].map({True: 'Active', False: 'Inactive'}).value_counts().to_dict()

        if request.format == ReportFormat.PDF:
            pdf_content = _create_pdf_report(
                title="Loan Distribution by Status",
                data={str(k): str(v) for k, v in status_counts.items()}
            )
            file_response = FileResponse(filename="status_distribution.pdf", content=pdf_content, mime_type="application/pdf")
            return LoanDistributionByStatusResponse(file_response=file_response)

        if request.format == ReportFormat.EXCEL:
            report_df = pd.DataFrame(list(status_counts.items()), columns=['Status', 'Count'])
            excel_content = _create_excel_report(report_df)
            file_response = FileResponse(
                filename="status_distribution.xlsx",
                content=excel_content,
                mime_type="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
            )
            return LoanDistributionByStatusResponse(file_response=file_response)
            
        if request.format == ReportFormat.PLOTLY_JSON:
            import plotly.express as px
            import json
            fig = px.pie(
                values=list(status_counts.values()), 
                names=list(status_counts.keys()), 
                title='Loan Distribution by Status'
            )
            plotly_json = fig.to_json()
            file_response = FileResponse(
                filename="status_distribution.json",
                content=plotly_json.encode('utf-8'),
                mime_type="application/json"
            )
            return LoanDistributionByStatusResponse(file_response=file_response)

        structured_response = LoanDistributionByStatusResponse.StructuredResponse(status_distribution=status_counts)
        return LoanDistributionByStatusResponse(structured_response=structured_response)

    async def GetLoanPortfolioSegmentation(self, request, context):
        pass  # TODO
