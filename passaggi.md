# Ebiten-Hello
Ho settato quest progetto per avere una referenza quando voglio creare
un'applicazione per windows, WSL e Android.

## Creare la libreria hello.aar
Per prima cosa bisogna settare le variabili di sistema:

    export ANDROID_HOME=$HOME/android/
	export ANDROID_SDK_ROOT=${ANDROID_HOME}/platforms/android-34
	export PATH=${ANDROID_HOME}/cmdline-tools/latest/bin:${ANDROID_HOME}/platform-tools:${PATH}
	export ANDROID_NDK_HOME=${ANDROID_HOME}/ndk/25.1.8937393

Poi si lancia ebitenmobile. Supponiamo che il clone della repository è in 

    ~/scratch/golang/ebiten_hello
 Per creare la libreria:   

    cd ~/scratch/golang/ebiten_hello/mobile
    ebitenmobile bind -target android -javapkg com.myapp.hello -o ./android/hello/hello.aar .
Nota che la directory in questione è la mobile e non android. Questo perché la stessa library
viene usata per il target IOS.

## Creare l'apk per Android
Per creare l'apk di Android occorre Gradle. 
Uso la stessa installazione e configurazione di AndroSolitario.
Qui bisogna settare le variabili di sistema in modo diverso:

    export ANDROID_SDK_ROOT=$HOME/android/
    export PATH=${ANDROID_SDK_ROOT}/cmdline-tools/latest/bin:${ANDROID_SDK_ROOT}/platform-tools:${PATH}
    cd ~/scratch/golang/ebiten_hello/mobile/android
	./gradlew compileDebugSources
    ./gradlew installDebug
    adb shell am start -n com.myapp.hello/.MainActivity

## Problemi
- Riconoscere lo smartphone in WSL2. 
- Far funzionare Gradle.