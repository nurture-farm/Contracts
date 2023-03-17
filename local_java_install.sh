#!/bin/bash
set -e

if [[ -z "$1" ]]
then
  echo "Please pass version as first argument example is 1.2.3 (major.minor.patch)"
fi

ARG_VERSION=$1

# Booking workflows Contracts
mvn install:install-file -Dfile=BookingServices/BookingWorkflows/Gen/JavaBookingServices/libbooking_workflows_proto-speed.jar -DgroupId=farm.nurture.core.contracts.workflows -DartifactId=booking_workflows_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=BookingServices/BookingWorkflows/Gen/JavaBookingServices/libbooking_workflows_java_grpc.jar -DgroupId=farm.nurture.core.contracts.workflows -DartifactId=booking_workflows -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# Supply workflows
mvn install:install-file -Dfile=BookingServices/SupplyWorkflows/Gen/JavaBookingServices/libsupply_workflows_proto-speed.jar -DgroupId=farm.nurture.core.contracts.workflows -DartifactId=supply_workflows_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=BookingServices/SupplyWorkflows/Gen/JavaBookingServices/libsupply_workflows_java_grpc.jar -DgroupId=farm.nurture.core.contracts.workflows -DartifactId=supply_workflows -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# Booking nexus contracts
mvn install:install-file -Dfile=BookingServices/BookingNexus/Gen/JavaBookingServices/libbooking_nexus_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=bn_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=BookingServices/BookingNexus/Gen/JavaBookingServices/libbooking_nexus_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=bn -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# SAM contracts
mvn install:install-file -Dfile=BookingServices/SupplyAvailabilityManager/Gen/JavaBookingServices/libsam_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=sam_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=BookingServices/SupplyAvailabilityManager/Gen/JavaBookingServices/libsam_enums_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=sam_enums -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=BookingServices/SupplyAvailabilityManager/Gen/JavaBookingServices/libsam_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=sam -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# PE Contracts
mvn install:install-file -Dfile=BookingServices/PlanningEngine/Gen/JavaBookingServices/libpe_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=pe_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=BookingServices/PlanningEngine/Gen/JavaBookingServices/libpe_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=pe -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# Rewards Contracts
mvn install:install-file -Dfile=BookingServices/RewardsService/Gen/JavaBookingServices/librewards_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=rewards_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=BookingServices/RewardsService/Gen/JavaBookingServices/librewards_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=rewards -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# Performance tracker  Contracts
mvn install:install-file -Dfile=PerformanceTracker/Gen/JavaPerformanceTracker/libperformance_tracker_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=communication_engine_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=PerformanceTracker/Gen/JavaPerformanceTracker/libperformance_tracker_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=communication_engine -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# Communication engine Contracts
mvn install:install-file -Dfile=CommunicationEngine/Gen/JavaCommunicationEngine/libcommunication_engine_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=communication_engine_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=CommunicationEngine/Gen/JavaCommunicationEngine/libcommunication_engine_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=communication_engine -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# Shield Contracts
mvn install:install-file -Dfile=Shield/Gen/JavaShield/libshield_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=shield_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=Shield/Gen/JavaShield/libshield_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=shield -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# Insights Contracts
mvn install:install-file -Dfile=BookingServices/BookingInsights/Gen/JavaBookingServices/libbooking_insights_proto-speed.jar -DgroupId=farm.nurture.core.contracts.insights -DartifactId=booking_insights_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=BookingServices/BookingInsights/Gen/JavaBookingServices/libbooking_insights_java_grpc.jar -DgroupId=farm.nurture.core.contracts.insights -DartifactId=booking_insights -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

# Common contracts
mvn install:install-file -Dfile=Common/Gen/JavaCommon/libevent_reference_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=event_reference -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=Common/Gen/JavaCommon/libtime_slot_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=time_slot -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=Common/Gen/JavaCommon/libcommon_enums_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=enums -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=Common/Gen/JavaCommon/libentities_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=entities -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=Common/Gen/JavaCommon/libactivities_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=activities -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=Common/Gen/JavaCommon/liberrors_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=errors -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=Common/Gen/JavaCommon/libtags_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=tags -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=Common/Gen/JavaCommon/libutils_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=utils -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

#loyalty enums
mvn install:install-file -Dfile=Common/Gen/JavaCommon/libloyalty_enums_proto-speed.jar  -DgroupId=farm.nurture.core.contracts.common -DartifactId=loyalty_action_enums -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;


#Loyalty Management
mvn install:install-file -Dfile=LoyaltyManagementSystem/LoyaltyManagementService/Gen/JavaLoyaltyManagementSystem/libloyalty_management_service_java_grpc.jar  -DgroupId=farm.nurture.core.contracts -DartifactId=loyalty -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=LoyaltyManagementSystem/LoyaltyManagementService/Gen/JavaLoyaltyManagementSystem/libloyalty_management_service_proto-speed.jar  -DgroupId=farm.nurture.core.contracts -DartifactId=loyalty_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

#Banking Jars
mvn install:install-file -Dfile=LoyaltyManagementSystem/BankingService/Gen/JavaLoyaltyManagementSystem/libbanking_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=banking -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=LoyaltyManagementSystem/BankingService/Gen/JavaLoyaltyManagementSystem/libbanking_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=banking_protos -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

#Mandi Service Jars
mvn install:install-file -Dfile=MandiService/Gen/JavaMandiService/libmandi_service_java_grpc.jar -DgroupId=farm.nurture.core.contracts -DartifactId=mandi_service -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;
mvn install:install-file -Dfile=MandiService/Gen/JavaMandiService/libmandi_service_proto-speed.jar -DgroupId=farm.nurture.core.contracts -DartifactId=mandi_service_proto -Dversion=${ARG_VERSION} -Dpackaging=jar -DgeneratePom=true;

