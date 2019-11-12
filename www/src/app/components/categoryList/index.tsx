import React from 'react';
import { Box, Button, Heading } from 'grommet';
import { Add } from 'grommet-icons';
import styled from 'styled-components';

interface CategoryListProps {
    className?: string;
    categories: string[];
}

const CategoryList: React.FC<CategoryListProps> = (props: CategoryListProps) => {
    const { className, categories } = props;

    return (
        <Box className={className} direction="column" width="150px">
            <Button icon={<Add size="small" />} label="Add" primary />
            <Heading level="4" margin="small">
                All topics
            </Heading>
            {categories.map((c: string) => (
                <Heading level="4" margin="small">
                    {c}
                </Heading>
            ))}
        </Box>
    );
};

export default styled(CategoryList)``;
