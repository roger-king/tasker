import React, { Suspense } from 'react';
import { Route, Switch, Redirect } from 'react-router-dom';
import { TasksPage, OverviewPage, SettingsPage, TaskPage, LoginPage } from './pages';

const RouterContainer: React.FC<{}> = (): JSX.Element => {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <Switch>
                <Route exact path="/tasker/admin/login" component={LoginPage} />
                <Route exact path="/tasker/admin" component={OverviewPage} />
                <Route exact path="/tasker/admin/tasks" component={TasksPage} />
                <Route exact path="/tasker/admin/settings" component={SettingsPage} />
                <Route exact path="/tasker/admin/task/:id" component={TaskPage} />
                <Redirect to="/tasker/admin" />
            </Switch>
        </Suspense>
    );
};

export default RouterContainer;
