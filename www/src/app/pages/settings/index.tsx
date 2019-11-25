import React from 'react';
import { Box, Heading, Button } from 'grommet';
import { Route, Switch, RouteComponentProps, Redirect } from 'react-router-dom';
import PluginSettingsPage from './plugin';

const SettingsPage: React.FC<RouteComponentProps> = (props: RouteComponentProps): JSX.Element => {
    const { match, history } = props;
    const { url } = match;

    return (
        <Box margin="xlarge">
            <Box margin={{ left: '200px' }}>
                <Heading> Settings </Heading>
            </Box>
            <Box direction="row" height="100%">
                <Box direction="column" width="150px" border={{ side: 'right', size: '3px' }} pad="medium">
                    <Button label="Plugin" plain onClick={() => history.push(`${url}/plugin`)} />
                </Box>
                <Switch>
                    <Route exact path={`${url}/plugin`} component={PluginSettingsPage} />
                    <Redirect from={`${url}`} to={`${url}/plugin`} />
                </Switch>
            </Box>
        </Box>
    );
};

export default SettingsPage;
