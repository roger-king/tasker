import React from 'react';
import { Heading, Image } from 'grommet';

const Logo: React.FC = (): JSX.Element => (
    <>
        <Image src={`${process.env.PUBLIC_URL}/images/flash.png`} width="40px" />
        <Heading level="2">TASKER.</Heading>
    </>
);

export default Logo;
