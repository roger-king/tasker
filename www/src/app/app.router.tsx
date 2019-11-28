import React, { useState, useEffect } from 'react';
import { Route, Redirect, RouteProps } from 'react-router-dom';
import { Box, Grid } from 'grommet';
import { TasksPage, OverviewPage, SettingsPage, TaskPage } from './pages';
import { check } from './data/auth';
import { REDIRECT_URL } from './app.constants';
import Header from './components/header';
import Notification from './components/notification';
import { UserContext } from './contexts/user';

export interface ProtectedRouteProps extends RouteProps {
    authenticationPath: string;
}

/* eslint-disable react/jsx-props-no-spreading */
export class ProtectedRoute extends Route<ProtectedRouteProps> {
    state = {
        isAuthenticated: true,
    };

    async componentDidMount(): Promise<void> {
        const isAuthenticated = await check();
        this.setState({ isAuthenticated });
    }

    public render(): JSX.Element {
        const { isAuthenticated } = this.state;
        let redirectPath = '';

        if (!isAuthenticated) {
            redirectPath = this.props.authenticationPath;
        }

        if (redirectPath) {
            sessionStorage.setItem(REDIRECT_URL, this.props.path as string);
            const renderComponent = (): JSX.Element => <Redirect to={{ pathname: redirectPath }} />;
            return <Route {...this.props} component={renderComponent} render={undefined} />;
        }
        return <Route {...this.props} />;
    }
}

const defaultProtectedRouteProps: ProtectedRouteProps = {
    authenticationPath: '/tasker/admin/login',
};

const RouterContainer: React.FC<{}> = (): JSX.Element => {
    /* eslint-disable-next-line */
    const [showNotification, setShowNotification] = useState<boolean>(false);
    const [UserStore, setUserStore] = useState<React.Context<User> | null>(null);

    useEffect(() => {
        UserContext().then(ctx => {
            setUserStore(ctx);
        });
    }, []);

    if (UserStore === null) {
        return <Box />;
    }

    return (
        <UserStore.Consumer>
            {(user: User) => (
                <Grid
                    fill
                    rows={['auto', 'flex']}
                    columns={['auto', 'flex']}
                    areas={[
                        { name: 'header', start: [0, 0], end: [1, 0] },
                        { name: 'main', start: [1, 1], end: [1, 1] },
                    ]}
                >
                    <Header gridArea="header" user={user} />
                    <Box gridArea="main" fill>
                        <ProtectedRoute
                            {...defaultProtectedRouteProps}
                            exact
                            path="/tasker/admin"
                            component={OverviewPage}
                        />
                        <ProtectedRoute
                            {...defaultProtectedRouteProps}
                            path="/tasker/admin/tasks"
                            component={TasksPage}
                        />
                        <ProtectedRoute
                            {...defaultProtectedRouteProps}
                            path="/tasker/admin/settings"
                            component={SettingsPage}
                        />
                        <ProtectedRoute
                            {...defaultProtectedRouteProps}
                            path="/tasker/admin/task/:id"
                            component={TaskPage}
                        />
                    </Box>
                    {showNotification && <Notification type="success" />}
                </Grid>
            )}
        </UserStore.Consumer>
    );
};

export default RouterContainer;
