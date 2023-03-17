BASE_DIR := $(shell pwd)
PROJECT_NAME := Contracts

.SILENT: start

.PHONY: all

all: clean package

clean:
	mvn clean

compile:
	mvn compile

package:
	mvn package

install:
	mvn clean install

install_sailor_local:
	bazel build //Sailor:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Sailor/Gen/JavaSailor
	cp -f -R ./bazel-bin/Sailor/libsailor_proto-speed.jar Sailor/Gen/JavaSailor
	cp -f -R ./bazel-bin/Sailor/libsailor_java_grpc.jar Sailor/Gen/JavaSailor
	mvn install:install-file -Dfile=Sailor/Gen/JavaSailor/libsailor_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=sailor_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true

install_mosaic_local:
	bazel build //Mosaic:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Mosaic/Gen/JavaMosaic
	cp -f -R ./bazel-bin/Mosaic/libmosaic_proto-speed.jar Mosaic/Gen/JavaMosaic
	cp -f -R ./bazel-bin/Mosaic/libmosaic_java_grpc.jar Mosaic/Gen/JavaMosaic
	mvn install:install-file -Dfile=Mosaic/Gen/JavaMosaic/libmosaic_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=mosaic_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true

install_galaxy_local:
	bazel build //Galaxy:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Galaxy/Gen/JavaGalaxn
	cp -f -R ./bazel-bin/Galaxy/libgalaxy_proto-speed.jar Galaxy/Gen/JavaGalaxy
	cp -f -R ./bazel-bin/Galaxy/libgalaxy_java_grpc.jar Galaxy/Gen/JavaGalaxy
	mvn install:install-file -Dfile=Galaxy/Gen/JavaGalaxy/libgalaxy_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=galaxy_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true

install_captain_local:
	bazel build //Captain:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Captain/Gen/JavaCaptain
	cp -f -R ./bazel-bin/Captain/libcaptain_proto-speed.jar Captain/Gen/JavaCaptain
	mvn install:install-file -Dfile=Captain/Gen/JavaCaptain/libcaptain_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=captain_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true

install_common_local:
	bazel build //Common:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Captain/Gen/JavaCommon
	cp -f -R ./bazel-bin/Common/libcommon_enums_proto-speed.jar Common/Gen/JavaCommon
	mvn install:install-file -Dfile=Common/Gen/JavaCommon/libcommon_enums_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=common_enums_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true

install_cms_local:
	bazel build //ContentManagementService:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Captain/Gen/JavaCMS
	cp -f -R ./bazel-bin/ContentManagementService/libcontent_management_service_proto-speed.jar ContentManagementService/Gen/JavaContentManagementService
	cp -f -R ./bazel-bin/ContentManagementService/libcontent_management_service_java_grpc.jar GamesService/Gen/JavaGamesService
	mvn install:install-file -Dfile=ContentManagementService/Gen/JavaContentManagementService/libcontent_management_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=content_management_service_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	mvn install:install-file -Dfile=ContentManagementService/Gen/JavaContentManagementService/libcontent_management_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=content_management_service_grpc -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit

force_clean:
	mvn dependency:purge-local-repository clean install

install_games_local:
	bazel build //GamesService:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Captain/Gen/JavaCMS
	cp -f -R ./bazel-bin/GamesService/libgames_service_java_grpc.jar GamesService/Gen/JavaGamesService
	cp -f -R ./bazel-bin/GamesService/libgames_service_proto-speed.jar GamesService/Gen/JavaGamesService
	mvn install:install-file -Dfile=GamesService/Gen/JavaGamesService/libgames_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=games_service_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	mvn install:install-file -Dfile=GamesService/Gen/JavaGamesService/libgames_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=games_service_grpc -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit

install_device_service_local:
	bazel build //DeviceService:all  --sandbox_debug --verbose_failures || exit
	cp -f -R ./bazel-bin/DeviceService/libdevice_service_java_grpc.jar DeviceService/Gen/JavaDeviceService
	cp -f -R ./bazel-bin/DeviceService/libdevice_service_proto-speed.jar DeviceService/Gen/JavaDeviceService
	mvn install:install-file -Dfile=DeviceService/Gen/JavaDeviceService/libdevice_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=device_service_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	mvn install:install-file -Dfile=DeviceService/Gen/JavaDeviceService/libdevice_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=device_service_grpc -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit

install_afs_local:
	bazel build //AFS:all  --sandbox_debug --verbose_failures || exit
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit

install_communication_engine_local:
	bazel build //CommunicationEngine:all  --sandbox_debug --verbose_failures || exit
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit

install_engagement_local:
	bazel build //EngagementService:all  --sandbox_debug --verbose_failures || exit
	cp -f -R ./bazel-bin/EngagementService/libengagement_service_java_grpc.jar EngagementService/Gen/JavaEngagementService
	cp -f -R ./bazel-bin/EngagementService/libengagement_service_proto-speed.jar EngagementService/Gen/JavaEngagementService
	mvn install:install-file -Dfile=EngagementService/Gen/JavaEngagementService/libengagement_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=engagement_service_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	mvn install:install-file -Dfile=EngagementService/Gen/JavaEngagementService/libengagement_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=engagement_service_grpc -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit

install_offer_local:
	bazel build //OfferService:all  --sandbox_debug --verbose_failures || exit
	cp -f -R ./bazel-bin/OfferService/liboffer_service_java_grpc.jar OfferService/Gen/JavaOfferService
	cp -f -R ./bazel-bin/OfferService/liboffer_service_proto-speed.jar OfferService/Gen/JavaOfferService
	mvn install:install-file -Dfile=OfferService/Gen/JavaOfferService/liboffer_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=offer_service_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	mvn install:install-file -Dfile=OfferService/Gen/JavaOfferService/liboffer_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=offer_service_grpc -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit

install_common_local:
	bazel build //GamesService:all  --sandbox_debug --verbose_failures || exit
	mkdir -p Captain/Gen/JavaCMS
	cp -f -R ./bazel-bin/GamesService/libgames_service_java_grpc.jar GamesService/Gen/JavaGamesService
	cp -f -R ./bazel-bin/GamesService/libgames_service_proto-speed.jar GamesService/Gen/JavaGamesService
	mvn install:install-file -Dfile=GamesService/Gen/JavaGamesService/libgames_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=games_service_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	mvn install:install-file -Dfile=GamesService/Gen/JavaGamesService/libgames_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=games_service_grpc -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit

install_hobbit_local:
	bazel build //Hobbit:all  --sandbox_debug --verbose_failures || exit
	cp -f -R ./bazel-bin/Hobbit/libhobbit_java_grpc.jar Hobbit/Gen/JavaHobbit
	cp -f -R ./bazel-bin/Hobbit/libhobbit_proto-speed.jar Hobbit/Gen/JavaHobbit
	mvn install:install-file -Dfile=Hobbit/Gen/JavaHobbit/libhobbit_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=hobbit_proto -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	mvn install:install-file -Dfile=Hobbit/Gen/JavaHobbit/libhobbit_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=hobbit_grpc -Dversion=latest -Dpackaging=jar -DgeneratePom=true
	/bin/bash  ./copy_cc_libs_and_sources.sh || exit
	/bin/bash  ./copy_go_libs_and_sources.sh || exit
	/bin/bash  ./copy_java_libs_and_sources.sh || exit
	/bin/bash  ./copy_python_libs_and_sources.sh || exit