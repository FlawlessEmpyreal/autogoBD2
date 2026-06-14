import os
import shutil
import random

src_images = r'F:\GoProject\file_bak\datasets\main\images'
src_labels = r'F:\GoProject\file_bak\datasets\main\labels'
dst_root   = r'F:\GoProject\file_bak\datasets\main'

for split in ['train', 'val']:
    os.makedirs(f'{dst_root}/images/{split}', exist_ok=True)
    os.makedirs(f'{dst_root}/labels/{split}', exist_ok=True)

images = [f for f in os.listdir(src_images) if f.endswith(('.jpg', '.png', '.jpeg'))]
random.shuffle(images)

split_idx  = int(len(images) * 0.8)
train_imgs = images[:split_idx]
val_imgs   = images[split_idx:]

for img in train_imgs:
    shutil.copy(f'{src_images}/{img}', f'{dst_root}/images/train/{img}')
    label = img.rsplit('.', 1)[0] + '.txt'
    if os.path.exists(f'{src_labels}/{label}'):
        shutil.copy(f'{src_labels}/{label}', f'{dst_root}/labels/train/{label}')

for img in val_imgs:
    shutil.copy(f'{src_images}/{img}', f'{dst_root}/images/val/{img}')
    label = img.rsplit('.', 1)[0] + '.txt'
    if os.path.exists(f'{src_labels}/{label}'):
        shutil.copy(f'{src_labels}/{label}', f'{dst_root}/labels/val/{label}')

print(f'训练集: {len(train_imgs)}张')
print(f'验证集: {len(val_imgs)}张')