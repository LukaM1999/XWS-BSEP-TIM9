package com.bsep.bsep.security;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.message.StringMapMessage;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.AuthenticationEntryPoint;
import org.springframework.stereotype.Component;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

@Component
public class RestAuthEntryPoint implements AuthenticationEntryPoint {

    private static final Logger logger = LogManager.getLogger("XML_ROLLING_FILE_APPENDER");

    @Override
    public void commence(HttpServletRequest request,
                         HttpServletResponse response,
                         AuthenticationException authException) throws IOException {
        StringMapMessage mapMessage = new StringMapMessage();
        mapMessage.put("msg", "Authentication failed");
        mapMessage.put("method", request.getMethod());
        mapMessage.put("url", request.getRequestURL().toString());
        String origin = request.getHeader("Origin");
        if (origin != null) {
            mapMessage.put("origin", origin);
        }
        String userAgent = request.getHeader("User-Agent");
        if (userAgent != null) {
            mapMessage.put("userAgent", userAgent);
        }

        logger.error(mapMessage);
        response.sendError(HttpServletResponse.SC_UNAUTHORIZED, authException.getMessage());
    }
}
