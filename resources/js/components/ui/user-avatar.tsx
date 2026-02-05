interface UserAvatarProps {
  /** User's full name */
  name?: string;
  /** Avatar size */
  size?: 'sm' | 'md' | 'lg';
  /** Image URL (optional) */
  src?: string;
  /** Additional classes */
  className?: string;
}

const sizeClasses = {
  sm: 'h-8 w-8 text-xs',
  md: 'h-10 w-10 text-sm',
  lg: 'h-12 w-12 text-base',
} as const;

function getInitials(name?: string): string {
  if (!name) return 'U';
  const parts = name.trim().split(' ');
  if (parts.length === 1) return parts[0].charAt(0).toUpperCase();
  return (parts[0].charAt(0) + parts[parts.length - 1].charAt(0)).toUpperCase();
}

export function UserAvatar({ name, size = 'md', src, className = '' }: UserAvatarProps) {
  const initials = getInitials(name);

  if (src) {
    return (
      <img
        src={src}
        alt={name || 'User avatar'}
        className={`rounded-full object-cover ${sizeClasses[size]} ${className}`}
      />
    );
  }

  return (
    <div
      className={`flex items-center justify-center rounded-full bg-[#1e3a8a] font-bold text-white ${sizeClasses[size]} ${className}`}
    >
      {initials}
    </div>
  );
}
