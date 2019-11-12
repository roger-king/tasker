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
        <Box gridArea={gridArea} height="80px" width="100vw" direction="row" align="center" gap="xsmall">
            <Box background="brand" width="300px" height="100%" align="center" justify="center">
                <Heading level="1">Tasker</Heading>
            </Box>
            <Box
                background="brand"
                align="center"
                justify="center"
                direction="row"
                width="100%"
                height="100%"
                pad="small"
            >
                <TextInput height="50px" style={{ width: '300px' }} />
                <DropButton icon={<User />} dropContent={<Box />} alignSelf="end" />
            </Box>
        </Box>
    );
};

export default styled(Header)``;
