# Install fastmap.
cd /home/zyh/GitHub/gocompile
sudo sh install.sh fastmap.tar.gz 

# Benchmark fastmap.
su - zyh -c 'cd /home/zyh/GitHub/gomapbench && benchok cmp && benchok rcmp'

# Install baseline.
cd /home/zyh/GitHub/gocompile
sudo sh install.sh baseline.tar.gz

# Benchmark baseline.
su - zyh -c 'cd /home/zyh/GitHub/gomapbench && benchok base && benchok rbase'
