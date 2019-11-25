import React, { useState } from 'react';
import { Box, Heading, Button } from 'grommet';
import { Route, Switch, RouteComponentProps, Redirect } from 'react-router-dom';
import PluginSettingsPage from './plugin';

const SettingsPage: React.FC<RouteComponentProps> = (props: RouteComponentProps): JSX.Element => {
    const { match, history } = props;
    const { url } = match;
    const [settingHeader, setSettingHeader] = useState<string>('Settings');

    return (
        <Box margin={{ left: '100px', right: '100px' }}>
            <Box margin={{ left: '200px' }}>
                <Heading> {settingHeader} </Heading>
            </Box>
            <Box direction="row" height="100%" margin={{ top: '40px' }}>
                <Box direction="column" width="150px" border={{ side: 'right', size: '3px' }} pad="medium">
                    <Button label="Plugin" plain onClick={() => history.push(`${url}/plugin`)} />
                </Box>
                <Box margin={{ left: '50px' }}>
                    <Switch>
                        <Route
                            exact
                            path={`${url}/plugin`}
                            render={() => {
                                setSettingHeader('Plugin Settings');
                                return <PluginSettingsPage />;
                            }}
                        />
                        <Redirect from={`${url}`} to={`${url}/plugin`} />
                    </Switch>
                </Box>
            </Box>
        </Box>
    );
};

export default SettingsPage;
