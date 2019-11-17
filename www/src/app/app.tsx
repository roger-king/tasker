import React, { Suspense } from 'react';
import { Grommet } from 'grommet';
import { BrowserRouter as Router, Switch, Route, Redirect } from 'react-router-dom';
import { LoginPage } from './pages';
import { theme } from './app.constants';
import RouterContainer from './app.router';

const App: React.FC = () => {
    return (
        <Grommet theme={theme} full>
            <Router>
                <Suspense fallback={<div>Loading...</div>}>
                    <Switch>
                        <Route exact path="/tasker/admin/login" component={LoginPage} />
                        <RouterContainer />
                        <Redirect to="/tasker/admin/login" />
                    </Switch>
                </Suspense>
            </Router>
        </Grommet>
    );
};

export default App;
