import React, { Suspense } from 'react';
import { Route, Switch, Redirect } from 'react-router-dom';
import { HomePage } from './pages';

const RouterContainer: React.FC<{}> = (): JSX.Element => {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <Switch>
                <Route exact path="/tasker/admin" component={HomePage} />
                <Redirect to="/tasker/admin" />
            </Switch>
        </Suspense>
    );
};

export default RouterContainer;
