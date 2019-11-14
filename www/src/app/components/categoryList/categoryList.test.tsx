import React from 'react';
import renderer from 'react-test-renderer';
import CategoryList from './index';

describe('<CategoryList />', () => {
    it('Matches the snapshot', () => {
        const select = jest.fn();
        const openModal = jest.fn();
        const list = renderer
            .create(
                <CategoryList
                    categories={['All', 'main', 'create']}
                    current={null}
                    selectCategory={select}
                    openModal={openModal}
                />,
            )
            .toJSON();
        expect(list).toMatchSnapshot();
    });
});
