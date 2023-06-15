FROM golang:latest
 
COPY demo_app .

# Expose application port
EXPOSE 3000
 
# Start app
CMD [ "./demo_app" ]