/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {}
    },
    plugins: [
        function ({ addVariant }) {
            addVariant('only-child', '&:only-child');
        },
    ],
};
