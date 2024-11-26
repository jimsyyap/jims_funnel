/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./src/**/*.{js,jsx,ts,tsx}",
    ],
    theme: {
        extend: {
            fontFamily: {
                // If you want to add custom fonts
                // 'sans': ['Inter', 'system-ui', 'sans-serif'],
            },
            colors: {
                // Custom color palette if needed
                // 'brand': {
                //   '50': '#f0f9ff',
                //   '100': '#e0f2fe',
                //   // Add more color shades
                // }
            },
        },
    },
    plugins: [],
}
