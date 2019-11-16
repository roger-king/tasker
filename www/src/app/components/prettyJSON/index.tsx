import React from 'react';
import { Box } from 'grommet';
import styled from 'styled-components';

interface PrettyJSONProps {
    className?: string;
    data: Record<string, any>;
    spaces?: number;
}

const PrettyJSON: React.FC<PrettyJSONProps> = (props: PrettyJSONProps) => {
    const { className, data, spaces = 4 } = props;

    return (
        <Box className={className} margin="small" pad="medium" background="light-4" align="center" justify="center">
            <pre>{JSON.stringify(data, null, spaces)}</pre>
        </Box>
    );
};

export default styled(React.memo(PrettyJSON))`
    min-width: 400px;
    max-width: 700px;
`;
