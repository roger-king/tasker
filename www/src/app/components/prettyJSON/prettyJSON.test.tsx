import React from 'react';
import Enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import renderer from 'react-test-renderer';
import PrettJSON from './index';

Enzyme.configure({ adapter: new Adapter() });

describe('<PrettyJSON />', () => {
    const data = {
        name: 'Steve Rogers',
        email: 'steve.rogers@avengers.com',
        friends: [
            {
                name: 'Tony Stark',
                alias: 'Iron Man',
            },
        ],
    };

    it('Matches the snapshot', () => {
        const prettyJSON = renderer.create(<PrettJSON data={data} />).toJSON();
        expect(prettyJSON).toMatchSnapshot();
    });
});
