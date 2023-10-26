create table if not exists collections
(
    id      uuid primary key default gen_random_uuid(),
    user_id uuid not null,
    name    text not null
);

create table if not exists documents
(
    id            uuid primary key default gen_random_uuid(),
    user_id       uuid not null,
    filename      text not null,
    path          text not null,
    collection_id uuid references collections (id) ON DELETE CASCADE
);

create table if not exists document_embeddings
(
    id          uuid primary key default gen_random_uuid(),
    document_id uuid references documents (id) ON DELETE CASCADE,
    page        integer not null,
    text        text    not null,
    embedding   vector(1536)
);

create table if not exists openai_usage
(
    id            uuid primary key   default gen_random_uuid(),
    user_id       uuid      not null,
    created_at    timestamp not null default now(),
    model         text      not null,
    input_tokens  int       not null,
    output_tokens int       not null
);

create table if not exists chat_message
(
    id            uuid primary key   default gen_random_uuid(),
    user_id       uuid      not null,
    created_at    timestamp not null default now(),
    collection_id uuid      not null references collections (id) ON DELETE CASCADE,
    prompt        text      not null,
    completion    text      not null
);

create table if not exists chat_message_source
(
    id                     uuid primary key default gen_random_uuid(),
    chat_message_id        uuid not null references chat_message (id) ON DELETE CASCADE,
    document_embeddings_id uuid not null references document_embeddings (id) ON DELETE CASCADE
);

create table if not exists payments
(
    id      uuid primary key   default gen_random_uuid(),
    date    timestamp not null default now(),
    user_id uuid      not null,
    amount  integer   not null
);