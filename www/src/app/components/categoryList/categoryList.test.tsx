import React from 'react';
import renderer from 'react-test-renderer';
import CategoryList from './index';

describe('<CategoryList />', () => {
    const categories = ['All', 'main', 'create'];
    let currentCategory: string | null = null;
    it('Matches the snapshot with ALL selected', () => {
        const select = jest.fn();
        const openModal = jest.fn();
        const list = renderer
            .create(
                <CategoryList
                    categories={categories}
                    current={currentCategory}
                    selectCategory={select}
                    openModal={openModal}
                />,
            )
            .toJSON();
        expect(list).toMatchSnapshot();
    });

    it('Matches the snapshot with MAIN selected', () => {
        currentCategory = categories[1];
        const select = jest.fn();
        const openModal = jest.fn();
        const list = renderer
            .create(
                <CategoryList
                    categories={categories}
                    current={currentCategory}
                    selectCategory={select}
                    openModal={openModal}
                />,
            )
            .toJSON();
        expect(list).toMatchSnapshot();
    });
});
