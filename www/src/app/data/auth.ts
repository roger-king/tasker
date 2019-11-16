import { defaultOAuthProvider } from '../app.constants';

export const authenticate = async (
    code: string,
): Promise<{ data: GithubOAuthLoginResponse | GithubOAuthErrorResponse; error?: string }> => {
    const data = await fetch(`/oauth/authenticate/${code}`, { method: 'POST' });
    return data.json();
};

export const fetchUserClientId = async (): Promise<{ data: OAuthProviderClientIdResponse }> => {
    let uri = '/oauth';
    if (defaultOAuthProvider === 'github') {
        uri += '/github/';
    } else {
        return { data: { client_id: '' } };
    }

    const data = await fetch(`${uri}/user`);
    return data.json();
};
