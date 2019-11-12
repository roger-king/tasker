import React from 'react';
import { Grommet, Grid, Box } from 'grommet';
import { BrowserRouter as Router } from 'react-router-dom';
import { theme } from './app.constants';
import Header from './components/header';
import RouterContainer from './app.router';

const App: React.FC = () => {
    return (
        <Grommet theme={theme} full>
            <Grid
                fill
                rows={['auto', 'flex']}
                columns={['auto', 'flex']}
                areas={[
                    { name: 'header', start: [0, 0], end: [1, 0] },
                    { name: 'main', start: [1, 1], end: [1, 1] },
                ]}
            >
                <Router>
                    <Header gridArea="header" />
                    <Box gridArea="main" fill>
                        <RouterContainer />
                    </Box>
                </Router>
            </Grid>
        </Grommet>
    );
};

export default App;
