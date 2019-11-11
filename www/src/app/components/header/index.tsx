import React from 'react';
import styled from 'styled-components';
import { Box, Heading, TextInput, DropButton } from 'grommet';
import { User } from 'grommet-icons';

interface HeaderProps {
    className?: string;
    gridArea: string;
}

const Header: React.FC<HeaderProps> = (props: HeaderProps): JSX.Element => {
    const { gridArea } = props;

    return (
        <Box
            gridArea={gridArea}
            height="80px"
            width="100vw"
            background="brand"
            direction="row"
            align="center"
            gap="small"
        >
            <Heading level="1">Tasker</Heading>
            <TextInput height="50px" style={{ width: '300px' }} />
            <DropButton icon={<User />} dropContent={<Box />} alignSelf="end" />
        </Box>
    );
};

export default styled(Header)``;
