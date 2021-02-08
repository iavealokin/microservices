FROM rabbitmq:3-management
ENV RABBITMQ_USER remote
ENV RABBITMQ_PASSWORD Cfyz11005310
ENV RABBITMQ_PID_FILE /var/lib/rabbitmq/mnesia/rabbitmq

ADD init.sh /init.sh
RUN chmod +x /init.sh
EXPOSE 15672

# Define default command
CMD ["/init.sh"]