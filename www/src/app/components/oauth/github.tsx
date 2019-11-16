import React, { useEffect } from 'react';
import { Button } from 'grommet';
import { Github } from 'grommet-icons';

interface GithubOAuthLoginBtnProps {
    clientId: string;
    scope: string[];
    isDisabled: boolean;
}

const GithubOAuthLoginBtn: React.FC<GithubOAuthLoginBtnProps> = (props: GithubOAuthLoginBtnProps) => {
    const { clientId, scope, isDisabled } = props;

    const stringOfScopes = scope.join('%20');

    useEffect(() => {
        const dom = document.getElementById('github-oauth-login-btn');
        if (dom && !isDisabled) {
            dom.setAttribute(
                'href',
                `https://github.com/login/oauth/authorize?client_id=${clientId}&scope=${stringOfScopes}`,
            );
        }
    }, [clientId, isDisabled, stringOfScopes]);

    return (
        <Button
            id="github-oauth-login-btn"
            href="/"
            icon={<Github size="medium" />}
            label="Github Login"
            style={{ borderRadius: '5px' }}
            primary
            disabled={isDisabled}
        />
    );
};

export default GithubOAuthLoginBtn;
