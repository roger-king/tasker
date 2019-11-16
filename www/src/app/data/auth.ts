import { defaultOAuthProvider } from '../app.constants';

export const authenticate = async (code: string): Promise<GithubOAuthLoginResponse> => {
    const data = await fetch(`/authenticate/${code}1`);
    return data.json();
};

export const fetchClientId = async (scope: string): Promise<{ data: OAuthProviderClientIdResponse }> => {
    let uri = '/oauth';
    if (defaultOAuthProvider === 'github') {
        uri += '/github/';
    } else {
        return { data: { client_id: '' } };
    }

    const data = await fetch(`${uri}/${scope}`);
    return data.json();
};
