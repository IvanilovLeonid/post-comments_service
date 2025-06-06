import { gql } from '@apollo/client';

export const GET_ALL_POSTS = gql`
    query CheckAllPosts {
        posts(first: 1) { # указывает какая страница
            edges {
                node {
                    id
                    title
                }
            }
        }
    }
`;

export const CREATE_POST = gql`
    mutation CreateFirstPost {
        createPost(input: {
            title: "Тестовый пост 1",
            content: "Содержание первого тестового поста",
            author: "Администратор",
            allowComments: true
        }) {
            id
        }
    }
`;
