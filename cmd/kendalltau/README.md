# Kendall's Tau

We can start the program on a SLURM cluster using the following command line:
```
srun -p K40x4  --cpus-per-task=20  -N 1 -n 1 ./kendalltau -p 20 /mnt/data/ssatheesh/models/evals/costs.npy
```
