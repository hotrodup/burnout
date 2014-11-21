FROM dockerfile/nodejs

ADD . /src
WORKDIR /src
RUN npm install

# Define default command.
CMD ["npm", "start"]
# Expose ports.
EXPOSE 5000
