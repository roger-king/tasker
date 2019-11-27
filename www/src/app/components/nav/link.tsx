import React from 'react';
import { Box } from 'grommet';
import { NavLink } from 'react-router-dom';
import styled from 'styled-components';

interface LinkProps {
    className?: string;
    to: string;
    name: string;
    disabled: boolean;
    icon?: JSX.Element;
}

const Link: React.FC<LinkProps> = (props: LinkProps) => {
    const { className, to, name } = props;

    return (
        <Box className={className} direction="row" align="center" gap="xsmall">
            <NavLink
                className="navbar-link"
                to={to}
                isActive={(_match, location): boolean => {
                    if (location.pathname === to) return true;

                    return false;
                }}
                activeStyle={{
                    width: '100%',
                    color: 'white',
                    fontWeight: 'bold',
                    background: 'rgba(19, 156, 219)',
                }}
            >
                {name}
            </NavLink>
        </Box>
    );
};

export default styled(Link)`
    .navbar-link {
        display: inline-block;
        text-align: center;
        margin:0 auto;
        height: 50px;
        width: 100%;
        margin-top: 10px;
        margin-right: 10px;
        pointer: ${(props: LinkProps): string => (props.disabled ? 'not-allowed' : 'pointer')}
        color:  'white';
        text-decoration: none;
    }
    .navbar-link:hover {
        background: #139cdb;
    }
`;
