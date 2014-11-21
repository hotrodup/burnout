FROM dockerfile/nodejs

# Set instructions on build.
ONBUILD ADD package.json /src/
ONBUILD RUN npm install
ONBUILD ADD . /src

# Define working directory.
WORKDIR /src

# Define default command.
CMD ["npm", "start"]

# Expose ports.
EXPOSE 5000