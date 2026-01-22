/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        background: {
          deep: '#0B0C10',
          dark: '#121212',
        },
        surface: {
          card: '#1C1C1E',
          widget: '#18181B',
          border: '#2E2E30',
        },
        accent: {
          neon: '#00FFC2',
          neonHover: '#00E6B0',
        },
        status: {
          error: '#FF4D4D',
          warning: '#FFC107',
          info: '#2196F3',
          success: '#4CAF50',
        },
        text: {
          primary: '#FFFFFF',
          secondary: '#A1A1AA',
          muted: '#52525B',
        }
      },
      borderRadius: {
        'xl': '12px',
        '2xl': '16px',
      },
      boxShadow: {
        'neon': '0 0 10px rgba(0, 255, 194, 0.4), 0 0 20px rgba(0, 255, 194, 0.2)',
        'neon-hover': '0 0 15px rgba(0, 255, 194, 0.6), 0 0 30px rgba(0, 255, 194, 0.3)',
      },
      gridTemplateColumns: {
        '24': 'repeat(24, minmax(0, 1fr))',
      },
      gridTemplateRows: {
        '7': 'repeat(7, minmax(0, 1fr))',
      },
    },
  },
  plugins: [],
}

