import { gql } from '@apollo/client';

export const GET_ALL_POSTS = gql`
  query CheckAllPosts {
    posts(first: 1) {
      edges {
        node {
          id
          title
        }
      }
    }
  }
`;

export const GET_POST_WITH_COMMENTS = gql`
  query GetPostWithComments($id: ID!) {
    post(id: $id) {
      id
      title
      content
      comments(first: 1) {
        id
        author
        text
        replies(first: 1) {
          id
          author
          text
        }
      }
    }
  }
`;