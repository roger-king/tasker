import React from 'react';
import { Button } from 'grommet';
import { Github } from 'grommet-icons';

interface GithubOAuthLoginBtnProps {
    clientId: string;
    scope: string[];
}

const GithubOAuthLoginBtn: React.FC<GithubOAuthLoginBtnProps> = (props: GithubOAuthLoginBtnProps) => {
    const { clientId, scope } = props;

    const stringOfScopes = scope.join('%20');

    return (
        <Button
            icon={<Github size="medium" />}
            href={`https://github.com/login/oauth/authorize?client_id=${clientId}&scope=${stringOfScopes}`}
            label="Github Login"
            style={{ borderRadius: '5px' }}
            primary
        />
    );
};

export default GithubOAuthLoginBtn;
