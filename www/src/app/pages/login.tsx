import React, { useEffect, useState } from 'react';
import { Box } from 'grommet';
import { useLocation } from 'react-router-dom';
import qs, { ParsedQuery } from 'query-string';
import GithubOAuthLoginBtn from '../components/oauth/github';

import { authenticate } from '../data/auth';
import { GITHUB_CLIENT_ID, GITHUB_LOGIN_SCOPE, LOGIN_STATUS } from '../app.constants';

const LoginPage: React.FC = () => {
    const location = useLocation();
    const [status, setStatus] = useState<LOGIN_STATUS>(LOGIN_STATUS.INITIAL);

    useEffect(() => {
        const queryString = location.search;
        const q: ParsedQuery = qs.parse(queryString);

        if (q && q.code) {
            setStatus(LOGIN_STATUS.LOADING);

            const code = q.code;
            authenticate(code as string).then((data: any) => {
                console.log('authenticated', data);
                setStatus(LOGIN_STATUS.SUCCESS);
            });
            console.log('Getting github access token');
        }
    }, [location.search]);

    if (status === LOGIN_STATUS.LOADING) {
        return <Box>logging in...</Box>;
    }

    return (
        <Box>
            <GithubOAuthLoginBtn clientId={GITHUB_CLIENT_ID} scope={GITHUB_LOGIN_SCOPE} />
        </Box>
    );
};

export default LoginPage;
