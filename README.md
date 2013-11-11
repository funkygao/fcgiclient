fcgiclient
==========

fastcgi client

Purpose
=======

A user request handler data flow is as follows:
    
      +----------+
      | browser  |
      +----------+
          |
      +----------+
      | nginx    |
      +----------+
          |
      +----------+
      | haproxy  |
      +----------+
          |
      +----------+
      | nginx    |
      |----------|
      | php-fpm  |
      +----------+

If HTTP Status non-200 returned, how do we identify the error location?

fcgiclient is a help, which first judge if it's php-fpm error.
