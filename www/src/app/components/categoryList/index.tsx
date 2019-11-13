import React from 'react';
import { Box, Button, Heading } from 'grommet';
import { Add } from 'grommet-icons';
import styled from 'styled-components';

interface CategoryListProps {
    className?: string;
    categories: string[];
    current: string | null;
    selectCategory: any;
    openModal: any;
}

const CategoryList: React.FC<CategoryListProps> = (props: CategoryListProps): JSX.Element => {
    const { className, categories, current, selectCategory, openModal } = props;
    const borderOpts = { side: 'left', size: '6px' };

    return (
        <Box className={className} direction="column" width="200px" gap="small">
            <Button icon={<Add size="small" />} label="Add" primary onClick={() => openModal()} />
            <Box
                style={{ cursor: 'pointer' }}
                border={current === null ? borderOpts : null}
                onClick={() => {
                    selectCategory(null);
                }}
            >
                <Heading level="4" margin="xsmall">
                    All topics
                </Heading>
            </Box>
            {categories.map(
                (c: string): JSX.Element => (
                    <Box
                        key={c}
                        style={{ cursor: 'pointer' }}
                        border={current === c ? borderOpts : null}
                        onClick={() => {
                            selectCategory(c);
                        }}
                    >
                        <Heading level="4" margin="xsmall">
                            {c}
                        </Heading>
                    </Box>
                ),
            )}
        </Box>
    );
};

export default styled(CategoryList)``;
