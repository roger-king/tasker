import React from 'react';
import { Grommet, Grid, Box } from 'grommet';
import { theme } from './app.constants';

import Header from './components/header';

const App: React.FC = () => {
    return (
        <Grommet theme={theme} full>
            <Grid
                fill
                rows={['auto', 'flex']}
                columns={['auto', 'flex']}
                areas={[
                    { name: 'header', start: [0, 0], end: [0, 0] },
                    { name: 'main', start: [1, 0], end: [1, 0] },
                ]}
            >
                <Header gridArea="header" />
                <Box gridArea="main" background="light-2" />
            </Grid>
        </Grommet>
    );
};

export default App;
