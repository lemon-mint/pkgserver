CREATE TYPE pkgtype AS ENUM ('go', 'rust', 'java', 'node');
CREATE TYPE vcstype AS ENUM ('bzr', 'fossil', 'git', 'hg', 'svn');

CREATE TABLE packages (
    id          BIGSERIAL PRIMARY KEY,
    pkg_name    TEXT      NOT NULL,
    pkg_type    pkgtype   NOT NULL,
    vcs         vcstype   NOT NULL,
    url         TEXT      NOT NULL,
    description TEXT      NOT NULL,

    UNIQUE(pkg_name)
);

CREATE INDEX packages_search_fulltext ON packages USING GIN(to_tsvector('english', description));
CREATE INDEX packages_name ON packages USING HASH(pkg_name);
