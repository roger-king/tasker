import React, { useState, useEffect } from 'react';
import styled from 'styled-components';
import { Box, Button, DropButton, Heading, TextInput, Text } from 'grommet';
import { User, Notification, Configure } from 'grommet-icons';
import { useHistory } from 'react-router';

interface HeaderProps {
    className?: string;
    gridArea: string;
}

const Header: React.FC<HeaderProps> = (props: HeaderProps): JSX.Element => {
    const { gridArea } = props;
    const [notifcations, setNotifications] = useState<any[]>([]);
    const history = useHistory();

    useEffect(() => {
        setNotifications([]);
    }, []);

    return (
        <Box gridArea={gridArea} height="80px" width="100vw" direction="row" align="center" gap="xsmall">
            <Box
                background="brand"
                width="300px"
                height="100%"
                align="center"
                justify="center"
                onClick={(): void => history.push('/tasker/admin')}
            >
                <Heading level="1">TASKER.</Heading>
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
                    <Button onClick={(): void => history.push('/tasker/admin')}>Overview</Button>
                    <Button onClick={(): void => history.push('/tasker/admin/tasks')}>Tasks</Button>
                </Box>
                <Box direction="row" alignSelf="end" align="center">
                    <TextInput height="50px" style={{ width: '300px' }} />
                    <Button icon={<Configure />} onClick={(): void => history.push('/tasker/admin/settings')} />
                    <DropButton
                        id="header-notification-btn"
                        icon={<Notification />}
                        dropAlign={{ top: 'bottom', right: 'left' }}
                        dropContent={
                            <Box
                                id="header-notification-drop-content"
                                width="200px"
                                height="200px"
                                align="center"
                                justify="center"
                            >
                                {notifcations.length > 0 ? (
                                    <Box> We have notifications! </Box>
                                ) : (
                                    <Box>
                                        <Text>
                                            <i>All Caught Up</i>
                                        </Text>
                                    </Box>
                                )}
                            </Box>
                        }
                    />
                    <DropButton icon={<User />} dropContent={<Box />} />
                </Box>
            </Box>
        </Box>
    );
};

export default styled(Header)``;
