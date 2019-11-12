import React from 'react';
import styled from 'styled-components';
import { Box, Button, DropButton, Heading, TextInput } from 'grommet';
import { User, Notification, Configure } from 'grommet-icons';

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
                justify="between"
                direction="row"
                width="100%"
                height="100%"
                pad="small"
            >
                <Box direction="row" gap="medium">
                    <Button>Overview</Button>
                    <Button>Tasks</Button>
                </Box>
                <Box direction="row" alignSelf="end" align="center">
                    <TextInput height="50px" style={{ width: '300px' }} />
                    <DropButton icon={<Configure />} dropContent={<Box />} />
                    <DropButton icon={<Notification />} dropContent={<Box />} />
                    <DropButton icon={<User />} dropContent={<Box />} />
                </Box>
            </Box>
        </Box>
    );
};

export default styled(Header)``;
