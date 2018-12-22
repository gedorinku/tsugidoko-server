FROM ruby:2.5.3-alpine3.8
LABEL maintainer="Ryota Egusa <gedorinku@yahoo.co.jp>"

WORKDIR /db
COPY Gemfile Gemfile.lock ./

RUN apk add build-base postgresql-dev
RUN bundle install

ENTRYPOINT [ "ridgepole" ]
CMD [ "--version" ]
