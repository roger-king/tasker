import { ThemeType } from 'grommet';

export const theme: ThemeType = {
    global: {
        breakpoints: {
            small: {
                value: 900,
            },
            medium: {
                value: 3000,
            },
        },
        colors: {
            background: '#1a1d21',
            brand: '#282c34',
            'accent-1': '#9400D3',
            'accent-2': '#EF424C',
            'accent-3': '#F8F0EE',
            focus: 'accent-1',
            warning: '#B33A3A',
        },
        font: {
            family: 'Raleway, sans-serif',
        },
    },
    heading: {
        font: {
            family: 'Viga, sans-serif',
            size: '4em',
        },
        level: {
            2: {
                font: {
                    family: 'Viga, sans-serif',
                },
            },
            3: {
                font: {
                    family: 'Viga, sans-serif',
                },
            },
            4: {
                font: {
                    family: 'Viga, sans-serif',
                },
            },
        },
    },
};
