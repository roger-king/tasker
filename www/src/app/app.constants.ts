import { ThemeType } from 'grommet';
import { css } from 'styled-components';

export const REDIRECT_URL = 'tasker_redirect_url';

export const theme: ThemeType = {
    global: {
        colors: {
            background: '#1a1d21',
            brand: '#282c34',
            'accent-1': '#139DDB',
            'accent-2': '#EF424C',
            'accent-3': '#F8F0EE',
            focus: 'accent-1',
            warning: '#B33A3A',
        },
        font: {
            family: 'Lato, sans-serif',
        },
    },
    button: {
        border: {
            radius: '7px',
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
                    family: 'Lato, sans-serif',
                },
            },
        },
    },
    checkBox: {
        toggle: {
            background: { light: 'toggle-accent' },
            color: {
                light: 'toggle-knob',
            },
            size: '36px',
            knob: {
                extend: `
                top: -7px;
                box-shadow: 0px 0px 2px 0px rgba(0,0,0,0.12),
                 0px 2px 2px 0px rgba(0,0,0,0.24);
              `,
            },
            extend: ({ checked }): string => `
              height: 14px;
              ${checked &&
                  css`
                      background-color: #2196f3;
                      border-color: #2196f3;
                  `}
            `,
        },
    },
};

type GithubProvider = 'github';
type OAuthProvider = GithubProvider;
// Default is a local testing client id
export const defaultOAuthProvider: OAuthProvider = process.env.REACT_APP_OAUTH_PROVIDER
    ? (process.env.REACT_APP_OAUTH_PROVIDER as OAuthProvider)
    : 'github';

export const GITHUB_LOGIN_SCOPE = ['user'];

export enum LOGIN_STATUS {
    AUTHENTICATED,
    UNAUTHENTICATED,
    INITIAL,
    LOADING,
    ERROR,
    SUCCESS,
}
