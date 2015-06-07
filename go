#!/bin/bash
rm -frv build
mkdir -pv build/out/modules/pronto build/out/dt_image
mkdir -pv build/system/etc/firmware/wlan/prima
export ARCH=arm
export CROSS_COMPILE=/home/serenitys/cm11/prebuilts/gcc/linux-x86/arm/arm-eabi-4.7/bin/arm-eabi-
export STRIP=/home/serenitys/cm11/prebuilts/gcc/linux-x86/arm/arm-eabi-4.7/bin/arm-eabi-strip
make carmilla-lefanu-a500fu_defconfig
clear
make CONFIG_NO_ERROR_ON_MISMATCH=y -j10 && make modules
./tools/dtbTool -o build/out/dt_image/boot.img-dtb -s 2048 -p ./scripts/dtc/ ./arch/arm/boot/dts/
cp arch/arm/boot/*zImage build/out/boot.img-zImage
cp drivers/staging/prima/firmware_bin/WCNSS_cfg.dat build/system/etc/firmware/wlan/prima/WCNSS_cfg.dat
cp drivers/staging/prima/firmware_bin/WCNSS_qcom_cfg.ini build/system/etc/firmware/wlan/prima/WCNSS_qcom_cfg.ini
cp drivers/staging/prima/firmware_bin/WCNSS_qcom_wlan_nv.bin build/system/etc/firmware/wlan/prima/WCNSS_qcom_wlan_nv.bin
find -type f -name *.ko -exec cp {} build/out/modules/ \;
ls -al build/out/modules/
cd build/out/modules/
mv wlan.ko pronto/pronto_wlan.ko
cd ../../../
ls -al build/out/modules/ build/out/modules/pronto/ build/out/dt_image/ build/out/
echo Done !
