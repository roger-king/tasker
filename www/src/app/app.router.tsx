import React, { Suspense } from 'react';
import { Route, Switch, Redirect } from 'react-router-dom';
import { HomePage } from './pages';

const RouterContainer: React.FC<{}> = (): JSX.Element => {
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <Switch>
                <Route exact path="/" component={HomePage} />
                <Redirect to="/" />
            </Switch>
        </Suspense>
    );
};

export default RouterContainer;
