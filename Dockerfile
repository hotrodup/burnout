FROM dockerfile/nodejs

WORKDIR /src

ONBUILD ADD package.json /src
ONBUILD RUN npm install
ONBUILD ADD . /src

# Define default command.
CMD ["npm", "start"]

# Expose ports.
EXPOSE 5000
