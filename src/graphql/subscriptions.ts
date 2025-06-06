import { gql } from '@apollo/client';

export const COMMENTS_SUBSCRIPTION = gql`
  subscription OnCommentAdded($postId: ID!) {
    commentAdded(postId: $postId) {
      id
      author
      text
      createdAt
    }
  }
`;