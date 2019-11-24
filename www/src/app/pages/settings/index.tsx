import React from 'react';
import { Box, Heading } from 'grommet';

const SettingsPage: React.FC<{}> = (): JSX.Element => {
    return (
        <Box margin={{ left: '300px', right: '300px' }}>
            <Heading> Settings </Heading>
        </Box>
    );
};

export default SettingsPage;
