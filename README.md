# blockchain-take-home

A simple blockchain application that allows users to create, update, and delete blog posts. Only contains one custom x/ module, `blog`.

## Getting started

Starting from the `blog` directory:

1. Install the chain binary with `make install`
2. Initialize the genesis files with `make init`
3. Start the chain with `blogd start --minimum-gas-prices=0stake`

## Tasks

1. The blog module does not show up in `blogd query` and `blogd tx` commands. Wire the blog module into the app so that it can be queried and transactions can be made. ✅
2. Whenever we create a new blog post, it overwrites an existing post. Fix this issue so that we can create multiple posts. ✅
3. Updating a post doesn't work as expected. Figure out why and fix it. ✅
4. Add a `created_at` and `last_updated_at` timestamp field to the blog post type, and update it with the current block time accordingly whenever a post is created or updated. ✅
5. Bonus: Implement a grant authorization mechanism where a post creator can allow other addresses to update or delete their post (e.g. Bob can update and delete Alice's blog posts). ✅

## Useful commands

- `rm -rf ~/blog/` - Remove the chain data

### Transactions

- `blogd tx blog create-post hello world --from alice --chain-id blog` - Create a new post
- `blogd tx blog update-post "Hello" "Cosmos" 1 --from alice --chain-id blog` - Update a post
- `blogd tx blog delete-post 1 --from alice  --chain-id blog` - Delete a post
- `blogd tx blog grant-auth 1 bob --from alice --chain-id blog` - Grant authorization to Bob for post 1
- `blogd tx blog revoke-auth 1 bob --from alice --chain-id blog` - Revoke authorization from Bob for post 1

## Queries

- `blogd q blog show-post 1` - Show a post
- `blogd q blog list-post` - List all posts
