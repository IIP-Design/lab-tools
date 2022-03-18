FROM amazonlinux:latest

LABEL Name=gpalab/wordpress Version=1.0.0

ARG context

# Define variables.
ENV HTTPD_CONF /etc/httpd/conf/httpd.conf
ENV FPM_POOL /etc/php-fpm.d/www.conf
ENV FPM_CONF /etc/php-fpm.conf
ENV PHP_CONF /etc/php.ini
ENV TZ America/New_York

# Install core dependencies (Apache, Git, Nano, Zip, etc.).
RUN yum update -y
RUN yum install -y httpd httpd-tools procps nano git tar zip unzip

# Install PHP 7.4
RUN amazon-linux-extras enable php7.4
RUN yum install -y php php-{cli,common,fpm,gd,json,mbstring,mcrypt,mysqlnd,pdo,xml,xmlrpc,zip}

RUN yum clean all

# Install Composer.
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/bin --filename=composer

# Colorize terminal.
RUN touch ~/.bashrc && \
  echo "export LS_OPTIONS='--color=auto'" >> ~/.bashrc && \
  echo 'eval "$(dircolors -b)"' >> ~/.bashrc && \
  echo 'alias ls="ls $LS_OPTIONS"' >> ~/.bashrc && \
  source ~/.bashrc

# Add deploy user.
RUN groupadd -g 1001 deploy && \
  useradd -r -u 1001 -g deploy deploy

# Customize PHP settings.
RUN sed -i 's/user = apache/user = deploy/g' $FPM_POOL && \
  sed -i 's/group = apache/user = deploy/g' $FPM_POOL && \
  sed -i 's/;listen.owner = nobody/listen.owner = deploy/g' $FPM_POOL && \
  sed -i 's/;listen.group = nobody/listen.group = deploy/g' $FPM_POOL && \
  sed -i 's/;listen.mode = 0660/listen.mode = 0660/g' $FPM_POOL && \
  sed -i 's/listen.acl_users = apache,nginx/;listen.acl_users = apache,nginx/g' $FPM_POOL && \
  # php-fpm.conf
  sed -i 's/error_log = \/var\/log\/php-fpm\/error.log/error_log = \/proc\/self\/fd\/2/g' $FPM_CONF && \
  sed -i 's/\/var\/log\/php-fpm\/www-error.log/\/proc\/self\/fd\/2/g' $FPM_CONF && \
  # php.ini
  sed -i 's/log_errors_max_len = 1024/log_errors_max_len = 0/g' $PHP_CONF && \
  sed -i 's/memory_limit = 128M/memory_limit = 3000M/g' $PHP_CONF && \
  sed -i 's/post_max_size = 8M/post_max_size = 2500M/g' $PHP_CONF && \
  sed -i 's/upload_max_filesize = 2M/upload_max_filesize = 2000M/g' $PHP_CONF && \
  sed -i 's/;date.timezone =/date.timezone = "America\/New_York"/g' $PHP_CONF

# Customize Apache settings.
RUN sed -i 's/var\/www\/html/www\/local\/sites-enabled\//g' $HTTPD_CONF && \
  sed -i 's/User apache/User deploy/g' $HTTPD_CONF && \
  sed -i 's/Group apache/Group deploy/g' $HTTPD_CONF

# Link error logs to docker logs.
RUN ln -sf /dev/stdout var/log/httpd/${context}_error.log

# Set timezone on instance.
RUN ln -snf /user/share/zoneinfo/$TZ /etc/localtime 

# Expose instance to external traffic on port 80.
EXPOSE 80

# Start container by enabling Apache and PHP.
CMD /usr/sbin/php-fpm -D; /usr/sbin/httpd -D FOREGROUND;