import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import Enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import renderer from 'react-test-renderer';
import Header from './index';

Enzyme.configure({ adapter: new Adapter() });

describe('<Header />', () => {
    const user: User = {
        name: 'Steve Rogers',
        email: 'steve.rogers@avengers.org',
        githubURL: 'https://github.com/steve-rogers',
        username: 'steve-rogers',
        bio: 'The first avenger',
    };

    it('Matches the snapshot', () => {
        const header = renderer
            .create(
                <MemoryRouter>
                    <Header gridArea="header" user={user} />
                </MemoryRouter>,
            )
            .toJSON();
        expect(header).toMatchSnapshot();
    });
});
