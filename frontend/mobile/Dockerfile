# Use an official OpenJDK as a base image
FROM openjdk:8-jdk

# Install dependencies and 32-bit libraries for Android SDK
RUN dpkg --add-architecture i386 && \
    apt-get update && \
    apt-get install -y \
    wget \
    unzip \
    xz-utils \
    libstdc++6:i386 \
    libz1:i386 \
    curl \
    git

# Set environment variables for Android SDK
ENV ANDROID_SDK_ROOT /sdk
ENV PATH ${PATH}:${ANDROID_SDK_ROOT}/tools:${ANDROID_SDK_ROOT}/tools/bin:${ANDROID_SDK_ROOT}/platform-tools

# Download and install the Android SDK
RUN wget -q "https://dl.google.com/android/repository/sdk-tools-linux-4333796.zip" -O sdk-tools-linux.zip && \
    mkdir -p ${ANDROID_SDK_ROOT} && \
    unzip sdk-tools-linux.zip -d ${ANDROID_SDK_ROOT} && \
    rm sdk-tools-linux.zip

# Accept Android SDK licenses and install necessary packages
RUN yes | sdkmanager --licenses && \
    sdkmanager "platform-tools" "platforms;android-29" "build-tools;29.0.3"

# Install Flutter
ENV FLUTTER_HOME /flutter
ENV PATH ${PATH}:${FLUTTER_HOME}/bin

RUN git clone -b stable https://github.com/flutter/flutter.git ${FLUTTER_HOME} && \
    ${FLUTTER_HOME}/bin/flutter doctor

# Pre-download Flutter dependencies
RUN flutter precache

# Set the working directory and copy project files
WORKDIR /app
COPY . /app

# Ensure Flutter dependencies are up-to-date
RUN flutter pub get

# Build the APK in release mode
RUN flutter build apk --release

# Copy the APK to a known location for easy retrieval
RUN cp /app/build/app/outputs/flutter-apk/app-release.apk /app/app-release.apk
