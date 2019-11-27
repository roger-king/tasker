import React, { useState } from 'react';
import { Box, Grid, Heading, Button } from 'grommet';
import { Route, Switch, RouteComponentProps, Redirect } from 'react-router-dom';
import PluginSettingsPage from './plugin';

const SettingsPage: React.FC<RouteComponentProps> = (props: RouteComponentProps): JSX.Element => {
    const { match, history } = props;
    const { url } = match;
    const [settingHeader, setSettingHeader] = useState<string>('Settings');

    return (
        <Box margin={{ left: '100px', right: '100px' }} height="100%">
            <Box margin={{ left: '200px' }}>
                <Heading> {settingHeader} </Heading>
            </Box>
            <Grid
                fill
                areas={[
                    { name: 'nav', start: [0, 0], end: [0, 0] },
                    { name: 'main', start: [1, 0], end: [1, 0] },
                ]}
                columns={['small', 'flex']}
                rows={['flex']}
                gap="small"
            >
                <Box
                    gridArea="nav"
                    direction="column"
                    height="98%"
                    width="80%"
                    border={{ side: 'right', size: '3px' }}
                    pad="medium"
                >
                    <Button label="Plugin" plain onClick={(): void => history.push(`${url}/plugin`)} />
                </Box>
                <Box gridArea="main" margin={{ left: '50px' }}>
                    <Switch>
                        <Route
                            exact
                            path={`${url}/plugin`}
                            render={(): JSX.Element => {
                                setSettingHeader('Plugin Settings');
                                return <PluginSettingsPage />;
                            }}
                        />
                        <Redirect from={`${url}`} to={`${url}/plugin`} />
                    </Switch>
                </Box>
            </Grid>
        </Box>
    );
};

export default SettingsPage;
