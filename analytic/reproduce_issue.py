import sys
import os

# Add the current directory to sys.path
sys.path.append(os.getcwd())
sys.path.append(r"C:\Users\Arsinenko\programming\flux_bank\analytic")

from google.protobuf.wrappers_pb2 import BoolValue
from api.generated.custom_types_pb2 import GetAllRequest

def reproduce():
    print("Attempting to reproduce the error...")
    request = GetAllRequest()
    is_desc = False
    
    try:
        # This is the line causing the error in customer_repository.py
        print("Executing: request.is_desc.value = BoolValue(value=is_desc)")
        request.is_desc.value = BoolValue(value=is_desc)
        print("Success! No error raised.")
    except TypeError as e:
        print(f"Caught expected TypeError: {e}")
    except Exception as e:
        print(f"Caught unexpected exception: {type(e).__name__}: {e}")

    print("\nVerifying the fix...")
    try:
        request = GetAllRequest()
        # The proposed fix
        print("Executing: request.is_desc.value = is_desc")
        request.is_desc.value = is_desc
        print(f"Fix works! value is: {request.is_desc.value}")
    except Exception as e:
        print(f"Fix failed with: {type(e).__name__}: {e}")

if __name__ == "__main__":
    reproduce()
