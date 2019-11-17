import React, { Suspense } from 'react';
import { Route, Switch, Redirect, RouteProps } from 'react-router-dom';
import { TasksPage, OverviewPage, SettingsPage, TaskPage, LoginPage } from './pages';
import { check } from './data/auth';

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
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <Switch>
                <Route exact path="/tasker/admin/login" component={LoginPage} />
                <ProtectedRoute {...defaultProtectedRouteProps} exact path="/tasker/admin" component={OverviewPage} />
                <ProtectedRoute {...defaultProtectedRouteProps} exact path="/tasker/admin" component={OverviewPage} />
                <ProtectedRoute
                    {...defaultProtectedRouteProps}
                    exact
                    path="/tasker/admin/tasks"
                    component={TasksPage}
                />
                <ProtectedRoute
                    {...defaultProtectedRouteProps}
                    exact
                    path="/tasker/admin/settings"
                    component={SettingsPage}
                />
                <ProtectedRoute
                    {...defaultProtectedRouteProps}
                    exact
                    path="/tasker/admin/task/:id"
                    component={TaskPage}
                />
                <Redirect to="/tasker/admin/login" />
            </Switch>
        </Suspense>
    );
};

export default RouterContainer;
