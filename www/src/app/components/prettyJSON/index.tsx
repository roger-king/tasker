import React from 'react';
import { Box } from 'grommet';

interface PrettyJSONProps {
    data: Record<string, any>;
}

const PrettyJSON: React.FC<PrettyJSONProps> = (props: PrettyJSONProps) => {
    const { data } = props;

    return (
        <Box margin="small" pad="medium" background="light-4" align="center" justify="center">
            <pre>{JSON.stringify(data)}</pre>
        </Box>
    );
};

export default React.memo(PrettyJSON);
